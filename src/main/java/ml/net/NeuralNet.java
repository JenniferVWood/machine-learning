package ml.net;

import ml.perceptron.Input;
import ml.perceptron.Vector;
import ml.sigmoid.SigmoidNeuron;

import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class NeuralNet {
    private final int numLayers;
    private final int[] sizes;

    private final double epsilon = 0.000001;
    private final double momentum = 0.7f;

    // for now... this is our expected output
    final double expectedOutputs[][] = { { 0 }, { 1 }, { 1 }, { 0 } };

    // the list of Nodes contains its own structure,
    // so, while the 2-d array isn't strictly necessary
    // it does make it easier to reason about.
    private List<List<Node>> nodes;


    public NeuralNet(int[] sizes) {
        this.sizes = sizes;
        this .numLayers = sizes.length;
        initRandomWeightsAndBiases();
    }


    private void initRandomWeightsAndBiases() {
        Random random = new Random();

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
                layer.add(new SigmoidNeuron(random.nextDouble()));
            }
            nodes.add(layer);
        }

        // now plug all the layers together with random weights.
        // (this is where the performance cost of OO starts to become obvious
        // in this context
        for (int l = 0; l <= nodes.size(); l++) {
            // wire each node to all the nodes in the next layer
            List<Node> layer = nodes.get(l);
            for (Node node : layer) {
                for (Node childNode : nodes.get(l+1)) {
                    childNode.addInput(node, random.nextDouble());
                }
            }
        }
    }

//    /**
//     *  (from the paper this is based on (http://neuralnetworksanddeeplearning.com/chap1.html):
//     *
//     * Train the neural network using mini-batch stochastic
//     * gradient descent.
//     *
//     * The training_data is a list of tuples (x, y) representing the training inputs and corresponding desired outputs.
//     * The variables epochs and mini_batch_size are what you'd expect - the number of epochs to train for,
//     * and the size of the mini-batches to use when sampling. eta is the learning rate, ηη. If the optional argument test_data is supplied,
//     * then the program will evaluate the network after each epoch of training, and print out partial progress.
//     * This is useful for tracking progress, but slows things down substantially.
//     */
//    public void SGD(List<List<Double>> trainingData, int numEpochs, int miniBatchSize, double learningRateETA,  boolean showProgress) {
///*
//      if test_data: n_test = len(test_data)
//        n = len(training_data)
//        for j in xrange(epochs):
//            random.shuffle(training_data)
//            mini_batches = [
//                training_data[k:k+mini_batch_size]
//                for k in xrange(0, n, mini_batch_size)]
//            for mini_batch in mini_batches:
//                self.update_mini_batch(mini_batch, eta)
//            if test_data:
//                print "Epoch {0}: {1} / {2}".format(
//                    j, self.evaluate(test_data), n_test)
//            else:
//                print "Epoch {0} complete".format(j)
// */
//        int nTest = 0;
//        if (showProgress) {
//            nTest = trainingData.size();
//        }
//
//        int n = trainingData.size();
//
//        for (int j = 0; j < trainingData.size(); j++) {
//            List<List<List<Double>>> miniBatches = new ArrayList<>(); // whoah
//
//            Collections.shuffle(trainingData);
//
//            for (int k = 0; k < n; k+= miniBatchSize) {
//                int end = k+miniBatchSize < trainingData.size()? k+miniBatchSize : trainingData.size() - 1;
//                List<List<Double>> miniBatch = trainingData.subList(k, end);
//                miniBatches.add(miniBatch);
//            }
//
//            for (List<List<Double>> miniBatch : miniBatches) {
//                updateMiniBatch(miniBatch, learningRateETA);
//            }
//        }
//
//    }
//


//    // this was a mess to unwrap using my conceptual model.  Abort!
//    private void updateMiniBatch(List<List<Double>> miniBatch, double learningRateETA) {
//        /*
//         def update_mini_batch(self, mini_batch, eta):
//        """Update the network's weights and biases by applying
//        gradient descent using backpropagation to a single mini batch.
//        The "mini_batch" is a list of tuples "(x, y)",
//                where x and y are lists of inputs and expected outputs
//        and "eta" is the learning rate."""
//
//        nabla_b = [np.zeros(b.shape) for b in self.biases]
//        nabla_w = [np.zeros(w.shape) for w in self.weights]
//        this just makes two new 2d arrays the size of the net, essentially
//
//        for x, y in mini_batch:
//            delta_nabla_b, delta_nabla_w = self.backprop(x, y)
//            nabla_b = [nb+dnb for nb, dnb in zip(nabla_b, delta_nabla_b)]
//            nabla_w = [nw+dnw for nw, dnw in zip(nabla_w, delta_nabla_w)]
//        self.weights = [w-(eta/len(mini_batch))*nw
//                        for w, nw in zip(self.weights, nabla_w)]
//        self.biases = [b-(eta/len(mini_batch))*nb
//                       for b, nb in zip(self.biases, nabla_b)]
//         */
//
//        Map<Node, List<Double>> nablas = new HashMap<>();
//        for (List<Double> x : miniBatch) {
//            for (Double y : x) {
////                nablas.add(backprop(x, y));
//            }
//        }
//    }


    // try using this implementation
    // https://kunuk.wordpress.com/2010/10/11/neural-network-backpropagation-with-java/

    private void runMiniBatch(List<List<Double>> inputs, List<List<Double>> expectedOutputs, double learningRateETA, int maxSteps) {
        double minError = 0.001; // for now... we should parameterize this.
//        List<List<Double>> resultOutputs = new ArrayList<>();
        int i;
        // Train neural network until minError reached or maxSteps exceeded
        double error = 1;
        for (i = 0; i < maxSteps && error > minError; i++) {
            error = 0;
            for (int p = 0; p < inputs.size(); p++) {
                setInputs(inputs.get(p));

                List<Double> output = getOutput();
//                resultOutputs.set(p, output);

                for (int j = 0; j < expectedOutputs.get(p).size(); j++) {
                    double err = Math.pow(output.get(j) - expectedOutputs.get(p).get(j), 2);
                    error += err;
                }

                applyBackpropagation(expectedOutputs.get(p), learningRateETA);
            }
        }

        printResult();

        System.out.println("Sum of squared errors = " + error);
        System.out.println("##### EPOCH " + i+"\n");
        if (i == maxSteps) {
            System.out.println("!Error training try again");
        } else {
//            printAllWeights();
//            printWeightUpdate();
        }
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

        for (Node n : getOutputLayer()) {
            List<Vector> vectors = n.getVectors();
            for (Vector vector : vectors) {
                double ak = n.value();
                double ai = vector.getInput().value();
                double desiredOutput = expectedOutput.get(i);

                double partialDerivative = -ak * (1 - ak) * ai
                        * (desiredOutput - ak);
                double deltaWeight = -learningRateETA * partialDerivative;
                double newWeight = vector.getWeight() + deltaWeight;
                vector.setDeltaWeight(deltaWeight);
                vector.setWeight(newWeight + momentum * vector.getPrevDeltaWeight());
            }
            i++;
        }

        for (List<Node> hiddenLayer : getHiddenLayers()) {
            // update weights for the hidden layers
            for (Node n : hiddenLayer) {
                List<Vector> vectors = n.getVectors();
                for (Vector con : vectors) {
                    double aj = n.value();
                    double ai = con.getInput().value();
                    double sumKoutputs = 0;
                    int j = 0;
                    for (Node out_neu : getOutputLayer()) {
                        double wjk = out_neu.getVectors().get(j).getWeight(); // not quite sure if "j" is correct param here. was "n.getId()"
                        double desiredOutput = expectedOutput.get(j);
                        double ak = out_neu.value();
                        j++;
                        sumKoutputs = sumKoutputs
                                + (-(desiredOutput - ak) * ak * (1 - ak) * wjk);
                    }

                    double partialDerivative = aj * (1 - aj) * ai * sumKoutputs;
                    double deltaWeight = -learningRateETA * partialDerivative;
                    double newWeight = con.getWeight() + deltaWeight;
                    con.setDeltaWeight(deltaWeight);
                    con.setWeight(newWeight + momentum * con.getPrevDeltaWeight());
                }
            }
        }
    }



    private void printResult()
    {
        System.out.println("NN example with xor training");
        for (int p = 0; p < getInputLayer().size(); p++) {
            System.out.print("INPUTS: ");
//            for (int x = 0; x < layers[0]; x++) {
//                System.out.print(inputs[p][x] + " ");
//            }
            for (Node node : getInputLayer()) {
                System.out.print(node.value() + " ");
            }

//            System.out.print("EXPECTED: ");
//            for (int x = 0; x < layers[2]; x++) {
//                System.out.print(expectedOutputs[p][x] + " ");
//            }

            for (int i = 0; i < getOutputLayer().size(); i++) {
                System.out.print(expectedOutputs[p][i] + " ");
            }


            System.out.print("ACTUAL: ");
            for (Node node : getOutputLayer()) {
                System.out.print(node.value() + " ");
            }
            System.out.println();
        }
        System.out.println();
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
