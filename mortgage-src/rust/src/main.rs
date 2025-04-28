mod calculator;

use calculator::MortgageCalculator;
use clap::Parser;
use std::fs::{self};
use std::io::{self};
use std::path::PathBuf;

/// CLI arguments struct
#[derive(Parser, Debug)]
#[command(author, version, about = "Mortgage payment calculator")]
struct Args {
    /// Source file containing mortgage data
    #[arg(short, long)]
    source: PathBuf,

    /// Optional target file for output (defaults to console)
    #[arg(short, long)]
    target: Option<PathBuf>,
}

fn main() -> io::Result<()> {
    let args = Args::parse();

    // Read input file contents
    let input_content = fs::read_to_string(&args.source)?;

    // Process the mortgage data
    let calculator = MortgageCalculator::new();
    let output_content = calculator.process(&input_content);

    // Output the results
    match &args.target {
        Some(path) => {
            // Write to the specified file
            fs::write(path, output_content)?;
        }
        None => {
            // Write to stdout
            print!("{}", output_content);
        }
    }

    Ok(())
}
