package ml.sigmoid;

import ml.perceptron.Perceptron;
import ml.perceptron.Vector;

import java.util.List;

public class SigmoidNeuron extends Perceptron {

    public SigmoidNeuron(double bias) {
        super(bias);
    }

    public double value() {
        return s(getVectors());
    }

    private double s(List<Vector> vectors) {
        double dot = super.fire();
        // somewhat goofy inheritance setup...
        double z = dot + super.getBias();

        return 1 / Math.pow(1 + Math.E, z);
    }
}
