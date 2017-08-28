package foo;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.*;

public class Threading {

    public static void main(String[] args) {
        one();
        two();
        three();
        four();
        five();
    }


    static void one() {
        Runnable task = () -> {
            String name = Thread.currentThread().getName();
            System.out.println("One " + name);
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println("One xx " + name);
        };

        task.run();
        Thread t = new Thread(task);
        t.start();
        System.out.println("done");
    }

    static void two() {
        ExecutorService executorService = Executors.newSingleThreadExecutor();
        executorService.submit(() ->{
            String name = Thread.currentThread().getName();
            System.out.println("Two " + name);
        });

        try {
            executorService.shutdown();
            executorService.awaitTermination(5, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            System.err.println("tasks interrupted");
        }finally {
            if (!executorService.isTerminated()) {
                System.err.println("cancel non-finished tasks");
            }
            executorService.shutdownNow();
            System.out.println("shutdown finished");
        }
    }

    static void three() {
        Callable<Double> doubleCallable = () -> {
            try {
                TimeUnit.SECONDS.sleep(2);
                return Math.PI;
            } catch(InterruptedException e) {
                throw new IllegalStateException("task interrupted", e);
            }
        };

        ExecutorService executorService = Executors.newFixedThreadPool(1);
        Future<Double> future = executorService.submit(doubleCallable);

        System.out.println("future done?" + future.isDone());
        try {
            Double result = future.get(1, TimeUnit.SECONDS);
            System.out.println("future done? " + future.isDone());
            System.out.print("result: " + result);
            executorService.shutdown();
        } catch (Exception e) {
            e.printStackTrace();
            executorService.shutdown();
        }

    }

    static void four() {
        List<Callable<String>> callables = Arrays.asList(
                () -> "c1"
                , () -> "c2"
                , () -> "c3"
        );
        ExecutorService executorService = Executors.newFixedThreadPool(1);
        try {
            executorService.invokeAll(callables)
                    .stream()
                    .map( future -> {
                        try {
                            return future.get();
                        } catch (Exception e) {
                            throw new RuntimeException(e);
                        }
                    })
                    .forEach(System.out::println);
            executorService.shutdown();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

    }


    static void five() {
        List<Callable<String>> callables = Arrays.asList(
                callable("task1", 3),
                callable("task2", 1),
                callable("task3", 3));

        ExecutorService executorService = Executors.newWorkStealingPool(4);
        try {
            String result = executorService.invokeAny(callables);
            System.out.println(result);
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            executorService.shutdown();
        }

    }

    static Callable<String> callable(String result, long sleepSeconds) {
        return () -> {
            TimeUnit.SECONDS.sleep(sleepSeconds);
            return result;
        };
    }
}
