package ml.perceptron;

/**
 * Describes relationship between a pair of nodes
 */
public class Vector {
    private double weight;
    private Perceptron input;

    public Vector(double weight, Perceptron input) {
        this.weight = weight;
        this.input = input;
    }

    public double getWeight() {
        return weight;
    }

    public void setWeight(double weight) {
        this.weight = weight;
    }

    public Perceptron getInput() {
        return input;
    }

    public void setInput(Perceptron input) {
        this.input = input;
    }
}
