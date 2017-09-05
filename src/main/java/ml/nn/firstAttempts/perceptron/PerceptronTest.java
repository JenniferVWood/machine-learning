package ml.nn.firstAttempts.perceptron;

public class PerceptronTest {

    public static void main(String[] args) {
        nandGate();

        adder();
    }

    public static void nandGate() {
        Input i1 = new Input(0.0);
        Input i2 = new Input(0.0);

        Perceptron p = new Perceptron(-3.0);
        p.addInput(i1, -2.0);
        p.addInput(i2, -2.0);

        System.out.println(p.value());
    }

    public static void adder() {
        Input i1 = new Input(1.0);
        Input i2 = new Input(1.0);

        Perceptron inputLayer = new Perceptron(-3.0);
        inputLayer.addInput(i1, -2.0);
        inputLayer.addInput(i2, -2.0);

        Perceptron hidden1 = new Perceptron(-3.0);
        hidden1.addInput(i1, -2.0);
        hidden1.addInput(inputLayer, -2.0);

        Perceptron hidden2 = new Perceptron(-3.0);
        hidden2.addInput(i2, -2.0);
        hidden2.addInput(inputLayer, -2.0);

        Perceptron oneDigitOut = new Perceptron(-3.0);
        oneDigitOut.addInput(hidden1, -2.0);
        oneDigitOut.addInput(hidden2, -2.0);

        Perceptron twoDigitOut = new Perceptron(-3.0);
        twoDigitOut.addInput(inputLayer, -2.0);
        twoDigitOut.addInput(inputLayer, -2.0);

        System.out.println("" + oneDigitOut.value() + "," + twoDigitOut.value());
    }
}
