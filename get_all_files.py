#!/usr/bin/env python3
import os
import sys

def main():
    # Define the output file name
    output_file = "combined.txt"
    
    # Open the output file for writing (will overwrite if exists)
    with open(output_file, "w", encoding="utf-8") as outfile:
        # Traverse the current directory and subdirectories
        for root, _, files in os.walk('.'):
            for file in files:
                if file.endswith(".go"):
                    filepath = os.path.join(root, file)
                    try:
                        with open(filepath, "r", encoding="utf-8", errors="replace") as infile:
                            # Write a comment line with the file's path at the top
                            outfile.write(f"// {filepath}\n")
                            # Write the content of the file
                            outfile.write(infile.read())
                            # Add a couple of newlines to separate files
                            outfile.write("\n\n")
                    except Exception as e:
                        print(f"Error reading {filepath}: {e}", file=sys.stderr)
    
    print(f"All .go files have been combined into {output_file}")

if __name__ == '__main__':
    main()
