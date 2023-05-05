public class Main {
    public static void main(String[] args) {
        System.out.println("Hello world!");
        Person personObj = new Person("Kuldeep",23,"Blue");
        String name = personObj.getName();
        int age = personObj.getAge();
        String color = personObj.getFavoriteColor();
        System.out.printf("my name is %s and I'm %d years old, my favorite color is %s\n",name,age,color);
    }
}