package ml.net;

import ml.perceptron.Perceptron;
import ml.perceptron.Vector;

import java.util.List;

public interface Node {
    double value();

    void addInput(Node input, double weight);

    List<Vector> getVectors();

    void setVectors(List<Vector> vectors);

    double getBias();

    void setBias(double bias);
}
