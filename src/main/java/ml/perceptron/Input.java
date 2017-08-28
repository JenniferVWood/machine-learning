package ml.perceptron;

/**
 * Special kind of Perceptron, used as input for a net:
 * value returns a literal, instead of a computed value.
 */
public class Input extends Perceptron {
    private double value;

    public Input(double value) {
        super(0.0);
        this.value = value;
    }

    @Override
    public double value() {
        return value;
    }
}
