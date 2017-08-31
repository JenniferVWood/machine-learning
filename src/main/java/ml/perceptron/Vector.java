package ml.perceptron;

import ml.net.Node;

/**
 * Describes relationship between a pair of nodes
 */
public class Vector {
    private double weight;
    private Node input;

    public Vector(double weight, Node input) {
        this.weight = weight;
        this.input = input;
    }

    public double getWeight() {
        return weight;
    }

    public void setWeight(double weight) {
        this.weight = weight;
    }

    public Node getInput() {
        return input;
    }

    public void setInput(Perceptron input) {
        this.input = input;
    }
}
