package ml.net;

import ml.perceptron.Input;
import ml.sigmoid.SigmoidNeuron;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class NeuralNet {
    private final int numLayers;
    private final int[] sizes;

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


}
