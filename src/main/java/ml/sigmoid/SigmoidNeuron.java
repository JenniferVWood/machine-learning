package ml.sigmoid;

import ml.perceptron.Perceptron;

public class SigmoidNeuron extends Perceptron {

    public SigmoidNeuron(double bias) {
        super(bias);
    }

    public double value() {
        return s();
    }

    private double s() {
        // somewhat goofy inheritance setup...
        double dot = super.fire();
        double z = dot + super.getBias();

        return 1 / Math.pow(1 + Math.E, z);
    }
}
