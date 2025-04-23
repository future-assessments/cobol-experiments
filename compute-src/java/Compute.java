public class Compute {


	public static void main(String[] args) {
		float fpl_pi = 3.14159265358979323f;
		float fpl_rad = 2.0f;
		float fpl_cir = fpl_pi * (2 * fpl_rad);
		System.out.printf("Perimeter %2.15f   Radius  %2.15f  fpl_pi %2.15f\n", fpl_cir, fpl_rad, fpl_pi);
	}
}
