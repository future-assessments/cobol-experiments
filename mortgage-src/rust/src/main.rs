mod calculator;

use calculator::MortgageCalculator;
use chrono::Datelike;
use clap::Parser;
use std::fs;
use std::io;
use std::path::PathBuf;

#[derive(Parser, Debug)]
#[command(author, version, about = "Mortgage payment calculator")]
struct Args {
    /// Source file containing mortgage data
    #[arg(short, long)]
    source: PathBuf,

    /// Optional target file for output (defaults to console)
    #[arg(short, long)]
    target: Option<PathBuf>,

    /// Starting year for mortgage calculations (defaults to current year)
    #[arg(short, long)]
    year: Option<u16>,
}

fn main() -> io::Result<()> {
    // Parse command line arguments
    let args = Args::parse();

    // Get the mortgage data from source file
    let input_content = fs::read_to_string(&args.source)?;

    // Get current year as default if year not specified
    let start_year = args.year.unwrap_or_else(|| chrono::Local::now().year() as u16);

    // Create calculator and process the mortgage data
    let calculator = MortgageCalculator::new();
    let output_content = calculator.process(&input_content, start_year);

    // Write output to target or stdout
    match args.target {
        Some(ref path) => {
            // Write to file
            fs::write(path, output_content)?;
        }
        None => {
            // Write to stdout
            print!("{}", output_content);
        }
    }

    Ok(())
}
