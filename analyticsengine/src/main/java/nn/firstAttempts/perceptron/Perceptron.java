package nn.firstAttempts.perceptron;

import nn.firstAttempts.Node;

import java.util.ArrayList;
import java.util.List;

public class Perceptron implements Node {
    private List<Vector> vectors = new ArrayList<>();
    private double bias = 0.0;

    public Perceptron(double bias) {
        this.bias = bias;
    }


    /**
     * recursively fires up the tree!
     *
     * It might be worth caching the output value, because
     * it will fire once for each child node, and the result won't change.
     *
     * Regardless, since doing it this way is a recursive, bottom-up tree traversal,
     * it's really slow.
     *
     * It's effective in that it shows what's happening *conceptually*, but we're
     * probably better off doing this with raw arrays, like it says in the book.
     */
    @Override
    public double value() {
        return fire() - bias >= 0? 1.0 : 0.0;
    }

    // in this simple implementation, just
    // give what I *think* is the dot product of inputs and weights.
    protected double fire() {
        double val = vectors.stream()
                .mapToDouble( v -> v.getInput().value() * v.getWeight() )
                .sum();

        return val;
    }

    @Override
    public void addInput(Node input, double weight) {
        Vector vector = new Vector(weight, input);
        vectors.add(vector);
    }

    @Override
    public List<Vector> getVectors() {
        return vectors;
    }

    @Override
    public void setVectors(List<Vector> vectors) {
        this.vectors = vectors;
    }

    @Override
    public double getBias() {
        return bias;
    }

    @Override
    public void setBias(double bias) {
        this.bias = bias;
    }
}
