import java.math.BigDecimal;

public class Compute2 {


	public static void main(String[] args) {
		var fpl_pi = new BigDecimal("3.14159265358979323");
		var fpl_rad = new BigDecimal("2.0");
		var multiplier = new BigDecimal("2");	
		var fpl_cir = fpl_pi.multiply(fpl_rad.multiply(multiplier));
		System.out.printf("Perimeter %2.15f   Radius  %2.15f  fpl_pi %2.15f\n", fpl_cir.doubleValue(), fpl_rad.doubleValue(), fpl_pi.doubleValue());
	}
}
