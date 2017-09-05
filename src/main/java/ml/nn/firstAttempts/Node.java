package ml.nn.firstAttempts;

import ml.nn.firstAttempts.perceptron.Vector;

import java.util.List;

public interface Node {
    double value();

    void addInput(Node input, double weight);

    List<Vector> getVectors();

    void setVectors(List<Vector> vectors);

    double getBias();

    void setBias(double bias);
}
