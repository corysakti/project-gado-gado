package self.cory.demokafka.common;

public class Sample {
    private String oneString;

    public Sample() {

    }

    public Sample(String oneString) {
        this.oneString = oneString;
    }

    public String getOneString() {
        return oneString;
    }

    public void setOneString(String oneString) {
        this.oneString = oneString;
    }

    @Override
    public String toString() {
        return "sample " + this.oneString + " ";
    }
}
