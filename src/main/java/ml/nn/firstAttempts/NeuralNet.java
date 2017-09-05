package ml.nn.firstAttempts;

import ml.nn.firstAttempts.perceptron.Vector;
import ml.nn.firstAttempts.perceptron.Input;
import ml.nn.firstAttempts.sigmoid.SigmoidNeuron;

import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class NeuralNet {
    private final int numLayers;
    private final int[] sizes;

    private final double epsilon = 0.001;
    private final double momentum = 0.7f;


    // the list of Nodes contains its own structure,
    // so, while the 2-d array isn't strictly necessary
    // it does make it easier to reason about.
    private List<List<Node>> nodes = new ArrayList<>();


    public static void main(String[] args) {

        // let's see if we can train it to build an xor circuit.
        NeuralNet net = new NeuralNet(new int[]{2, 4, 1});
        List<List<Double>> testInputs = new ArrayList<>();
        double i1[] = new double[] {1, 1};
        double i2[] = new double[] {1, 0};
        double i3[] = new double[] {0, 1};
        double i4[] = new double[] {0, 0};

        testInputs.add(toDList(i1));
        testInputs.add(toDList(i2));
        testInputs.add(toDList(i3));
        testInputs.add(toDList(i4));


        List<List<Double>> expectedOutput = new ArrayList<>();
        expectedOutput.add(toDList(new double[]{0.0}));
        expectedOutput.add(toDList(new double[]{1.0}));
        expectedOutput.add(toDList(new double[]{1.0}));
        expectedOutput.add(toDList(new double[]{0.0}));

        net.run(testInputs, expectedOutput, 0.9, 50000);
    }


    private static List<Double> toDList(double values[]) {
        return Arrays.stream(values).boxed().collect(Collectors.toList());
    }


    public NeuralNet(int[] sizes) {
        this.sizes = sizes;
        this .numLayers = sizes.length;
        initRandomWeightsAndBiases();
    }


    private double getRandom() {
        Random random = new Random();
        return (random.nextDouble() * 2 - 1);
    }

    private void initRandomWeightsAndBiases() {

        // do we need to define this as a separate instance variable?
        // I don't think so... yet...
        List<Node> inputs = new ArrayList<>();
        for (int i = 0; i < sizes[0]; i++) {
            Input input = new Input(0.0); // we'll set this later
            inputs.add(input);
        }
        nodes.add(inputs);

        // initialize the biases for the rest of the layers
        for (int l = 1; l < numLayers; l++) {
            List<Node> layer = new ArrayList<>();
            for (int n = 0; n < sizes[l]; n++) {

                layer.add(new SigmoidNeuron(getRandom()));
            }
            nodes.add(layer);
        }

        // now plug all the layers together with random weights.
        // (this is where the performance cost of OO starts to become obvious
        // in this context
        for (int l = 0; l < nodes.size() -1; l++) {
            // wire each node to all the nodes in the next layer
            List<Node> layer = nodes.get(l);
            for (Node node : layer) {
                for (Node childNode : nodes.get(l+1)) {
                    childNode.addInput(node, getRandom());
                }
            }
        }
    }

    /* I first attempted to do gradient descent from the code from (http://neuralnetworksanddeeplearning.com/chap1.html):
    * but that was a mess to translate, given my conceptual model.
    *
    * Using this implementation instead:
    * https://kunuk.wordpress.com/2010/10/11/neural-network-backpropagation-with-java/
    */
    private void run(List<List<Double>> testInputs, List<List<Double>> expectedOutputs, double learningRateETA, int maxSteps) {
        double minError = 0.001; // for now... we should parameterize this.

        List<List<Double>> actualOutputs = new ArrayList<>();

        int i;
        // Train neural network until minError reached or maxSteps exceeded
        double error = 1;
        for (i = 0; i < maxSteps && error > minError; i++) {
            error = 0;
            for (int p = 0; p < testInputs.size(); p++) {
                setInputs(testInputs.get(p));

                List<Double> output = getOutput();
                actualOutputs.add(output); // for reporting

                for (int j = 0; j < expectedOutputs.get(p).size(); j++) {
                    double err = Math.pow(output.get(j) - expectedOutputs.get(p).get(j), 2);
                    error += err;
                }

                applyBackpropagation(expectedOutputs.get(p), learningRateETA);
            }

            System.out.println("Sum of squared errors = " + error);
            System.out.println("##### EPOCH " + i+"\n");
        }


        if (i == maxSteps) {
            System.out.println("!Error training try again");
        } else {
//            printAllWeights();
//            printWeightUpdate();
        }
        printResult(testInputs, expectedOutputs, actualOutputs);
    }


    public void applyBackpropagation(List<Double> expectedOutput, double learningRateETA) {

        // error check, normalize value ]0;1[
        for (int i = 0; i < expectedOutput.size(); i++) {
            double d = expectedOutput.get(i);
            if (d < 0 || d > 1) {
                if (d < 0)
                    expectedOutput.set(i,0 + epsilon);
                else
                    expectedOutput.set(i, 1 - epsilon);
            }
        }

        int i = 0;


        // adjust output layer
        for (Node n : getOutputLayer()) {
            // weights
            List<ml.nn.firstAttempts.perceptron.Vector> vectors = n.getVectors();
            for (ml.nn.firstAttempts.perceptron.Vector vector : vectors) {
                double ak = n.value();
                double ai = vector.getInput().value();
                double desiredOutput = expectedOutput.get(i);

                double partialDerivative = -ak * (1 - ak) * ai
                        * (desiredOutput - ak);
                double deltaWeight = -learningRateETA * partialDerivative;
                double newWeight = vector.getWeight() + deltaWeight;
                vector.setDeltaWeight(deltaWeight);
                vector.setWeight(newWeight + momentum * vector.getPrevDeltaWeight());

                // for node
                double newBias = n.getBias() + deltaWeight;
                ((SigmoidNeuron) n).setDeltaBias(deltaWeight);
                n.setBias(newBias + momentum * ((SigmoidNeuron) n).getDeltaBias());

            }
            i++;
        }

        for (List<Node> hiddenLayer : getHiddenLayers()) {
            // update weights for the hidden layers
            for (Node hiddenNode : hiddenLayer) {
                List<ml.nn.firstAttempts.perceptron.Vector> vectors = hiddenNode.getVectors();

                // for each Vector in this Node
                for (ml.nn.firstAttempts.perceptron.Vector vector : vectors) {
                    double aj = hiddenNode.value();
                    double ai = vector.getInput().value();
                    double sumKoutputs = 0;
                    int vectorIndex = 0;

                    // find output Node
                    for (Node outNode : getOutputLayer()) {
                        // we want the weight for the vector from the current hidden Node to this output Node
//                        double wjk = outNode.getVectors().get(vectorIndex).getWeight(); // not quite sure if "j" is correct param here. was "n.getId()"
                        double wjk = getWeightForVector(hiddenNode, outNode);
                        double desiredOutput = expectedOutput.get(vectorIndex);
                        double ak = outNode.value();
                        vectorIndex++;
                        sumKoutputs = sumKoutputs
                                + (-(desiredOutput - ak) * ak * (1 - ak) * wjk);
                    }

                    double partialDerivative = aj * (1 - aj) * ai * sumKoutputs;
                    double delta = -learningRateETA * partialDerivative;
                    double newWeight = vector.getWeight() + delta;
                    vector.setDeltaWeight(delta);
                    vector.setWeight(newWeight + momentum * vector.getPrevDeltaWeight());

                    // for node
                    double newBias = hiddenNode.getBias() + delta;
                    ((SigmoidNeuron) hiddenNode).setDeltaBias(delta);
                    hiddenNode.setBias(newBias + momentum * ((SigmoidNeuron) hiddenNode).getDeltaBias());
                }
            }
        }
    }

    private Double getWeightForVector(Node from, Node to) {
        double ret = 0.0;
        for (Vector v : to.getVectors()) {
            if (v.getInput() == from) {
                return v.getWeight();
            }
        }
        throw new RuntimeException("Nodes are not connected.");
    }

    private void printResult(List<List<Double>> testInputs, List<List<Double>> expectedOutputs, List<List<Double>> resultOutputs)
    {
        System.out.println("NN example with xor training");

        // for each test set, print the expected and actual values
        for (int i = 0; i < testInputs.size(); i++) {
            System.out.print(
                    "INPUTS: ");
            System.out.print(testInputs.get(i)
                    .stream()
                    .map(d -> d + "")
                    .collect(Collectors.joining(" ")));

            System.out.print(" EXPECTED OUTPUTS: ");
            System.out.print(expectedOutputs.get(i)
                    .stream()
                    .map(d -> d + "")
                    .collect(Collectors.joining(" ")));

            System.out.print(" ACTUAL: ");
            System.out.print(resultOutputs.get(i)
                    .stream()
                    .map(d -> d + "")
                    .collect(Collectors.joining(" ")));

            System.out.println();
        }


//        for (int p = 0; p < getInputLayer().size(); p++) {
//            System.out.print("INPUTS: ");
//            for (Node node : getInputLayer()) {
//                System.out.print(node.value() + " ");
//            }


//            System.out.print("EXPECTED OUTPUTS: ");
//            for (int i = 0; i < getOutputLayer().size(); i++) {
//                System.out.print(expectedOutputs.get(p).get(i) + " ");
//            }


//            System.out.print("ACTUAL: ");
//            for (Node node : getOutputLayer()) {
//                System.out.print(node.value() + " ");
//            }
//            System.out.println();
    }



    private void setInputs(List<Double> inputs) {
        List<Node> newInputLayer = new ArrayList<>();
        for (Double d : inputs) {
            newInputLayer.add(new Input(d));
        }
        this.nodes.set(0, newInputLayer);
    }


    private List<Double> getOutput() {
        return nodes.get(nodes.size() -1)
                .stream()
                .mapToDouble(Node::value)
                .boxed()
                .collect(Collectors.toList());
    }


    private List<Node> getInputLayer() {
        return nodes.get(0);
    }

    private List<List<Node>> getHiddenLayers() {
        List<List<Node>> hiddenLayers = new ArrayList<>();
        for (int i = 1; i < nodes.size() - 2; i++) {
            hiddenLayers.add(nodes.get(i));
        }
        return hiddenLayers;
    }

    private List<Node> getOutputLayer() {
        return nodes.get(nodes.size() -1);
    }

    private List<Integer> evenNumbers() {
        return IntStream.range(0, 100)
                .filter(n -> n%2 == 0)
                .boxed()
                .collect(Collectors.toList());
    }
}
