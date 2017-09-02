package ml.perceptron;

import ml.net.Node;

/**
 * Describes relationship between a pair of nodes (double-linked list)
 */
public class Vector {
    private double weight;

    // for learning via backpropagation
    private double prevDeltaWeight;
    private double deltaWeight;

    private Node input;
    private Node output;

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

    public double getPrevDeltaWeight() {
        return prevDeltaWeight;
    }

    public void setPrevDeltaWeight(double prevDeltaWeight) {
        this.prevDeltaWeight = prevDeltaWeight;
    }

    public double getDeltaWeight() {
        return deltaWeight;
    }

    public void setDeltaWeight(double deltaWeight) {
        this.prevDeltaWeight = this.deltaWeight;
        this.deltaWeight = deltaWeight;
    }

    public void setInput(Node input) {
        this.input = input;
    }

    public Node getOutput() {
        return output;
    }

    public void setOutput(Node output) {
        this.output = output;
    }
}
