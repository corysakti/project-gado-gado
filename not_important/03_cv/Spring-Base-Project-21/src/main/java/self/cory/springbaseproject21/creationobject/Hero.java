package self.cory.springbaseproject21.creationobject;

import java.util.Objects;

public class Hero implements Comparable<Hero> {
    private String hp;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Hero hero = (Hero) o;

        return Objects.equals(hp, hero.hp);
    }

    @Override
    public int hashCode() {
        return hp != null ? hp.hashCode() : 0;
    }

    @Override
    protected Object clone() throws CloneNotSupportedException {
        return super.clone();
    }

    @Override
    public String toString() {
        return "Hero{" +
                "hp='" + hp + '\'' +
                '}';
    }

    @Override
    public int compareTo(Hero o) {
        return 0;
    }
}
