package ml.net;

import ml.perceptron.Perceptron;
import ml.sigmoid.SigmoidNeuron;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;
import java.util.stream.IntStream;

public class NeuralNet {
    private final int numLayers;
    private final int[] sizes;

    private List<List<Node>> nodes;


    public NeuralNet(int[] sizes) {
        this.sizes = sizes;
        this .numLayers = sizes.length;
        initRandomWeightsAndBiases();
    }

    private void initRandomWeightsAndBiases() {
        Random random = new Random();

        // initialize the biases
        for (int l = 0; l < numLayers; l++) {
            List<Node> layer = new ArrayList<>();
            for (int n = 0; n < sizes[l]; n++) {
                layer.add(new SigmoidNeuron(random.nextDouble()));
            }
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
