package ml.perceptron;

import java.util.ArrayList;
import java.util.List;

public class Perceptron {
    private List<Vector> vectors = new ArrayList<>();
    private double bias = 0.0;

    public Perceptron(double bias) {
        this.bias = bias;
    }


    public double value() {
        return fire() - bias >= 0? 1.0 : 0.0;
    }

    // in this simple implementation,
    protected double fire() {
        double val = 0.0;
        val = vectors.stream()
                .mapToDouble( v -> v.getInput().value() * v.getWeight() )
                .sum();

        return val;
    }

    public void addInput(Perceptron input, double weight) {
        Vector vector = new Vector(weight, input);
        vectors.add(vector);
    }

    public List<Vector> getVectors() {
        return vectors;
    }

    public void setVectors(List<Vector> vectors) {
        this.vectors = vectors;
    }

    public double getBias() {
        return bias;
    }

    public void setBias(double bias) {
        this.bias = bias;
    }
}
