fn main() {
    println!("Hello, compute!");
    let fpl_pi: f64 = 3.14159265358979323;
    let fpl_rad: f64 = 2.0;

    let fpl_cir: f64 = fpl_pi * (f64::from(2) * fpl_rad);

    println!("Perimieter {:.15}  Radius {:.16}    Pi {:.16}", fpl_cir, fpl_rad, fpl_pi);
}
