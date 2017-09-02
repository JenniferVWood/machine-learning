package ml.sigmoid;

import ml.perceptron.Perceptron;

public class SigmoidNeuron extends Perceptron {
    private double prevDeltaBias;
    private double deltaBias;

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

    public double getPrevDeltaBias() {
        return prevDeltaBias;
    }

    public void setPrevDeltaBias(double prevDeltaBias) {
        this.prevDeltaBias = prevDeltaBias;
    }

    public double getDeltaBias() {
        return deltaBias;
    }

    public void setDeltaBias(double deltaBias) {
        this.prevDeltaBias = this.deltaBias;
        this.deltaBias = deltaBias;
    }
}
