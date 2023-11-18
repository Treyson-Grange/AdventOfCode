import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
//I couldn't find my python version of this code, and this is HELLA unfinished so ignore this day
public class FindHighestElf {

    public int findGreates(String filename) throws FileNotFoundException {
        File file = new File(filename);
        Scanner scanner = new Scanner(file);

    }
    public static void main(String[] args) throws FileNotFoundException {
        FindHighestElf elf = new FindHighestElf();
        System.out.println(elf.findGreates("src/inputday1"));
    }
}
