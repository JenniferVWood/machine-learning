package ml.sigmoid;

import ml.perceptron.Input;
import ml.perceptron.Perceptron;

public class SigmoidNeuronTest {
    public static void main(String[] args) {
//        testOneNeuron();
        testOnePerceptron();
    }

    private static void testOneNeuron() {
        SigmoidNeuron neuron = new SigmoidNeuron(1.0);


        for (double value = -2.0; value < 2.0; value += 0.1) {
            for (double weight = -2.0; weight < 2.0; weight += 0.1) {
                Input input1 = new Input(value);
//                Input input2 = new Input(value -0.1);

                neuron.addInput(input1, weight);
//                neuron.addInput(input2, weight);

                System.out.println(/*"" + (value - 0.1) + " " +*/ value + " " + weight + ": " + neuron.value());
            }
        }
    }

    private static void testOnePerceptron() {
        Perceptron neuron = new Perceptron(1.0);

        for (double value = -2.1; value < 2.0; value += 0.1) {
            for (double weight = -2.0; weight < 2.0; weight += 0.1) {
                Input input1 = new Input(value);
                neuron.addInput(input1, weight);
                System.out.println(value + " " + weight + ": " + neuron.value());
                //........
            }
        }
    }
}
