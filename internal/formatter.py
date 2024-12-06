import json
import argparse

parser = argparse.ArgumentParser(description="Convert a file into JSON")
parser.add_argument("input_file", help="Path to input file")
parser.add_argument("output_file", help="Path to output file")

args = parser.parse_args()

with open(args.input_file, "r") as file:
    colum1 = []
    colum2 = []

    for line in file:
        values = line.split()
        colum1.append(int(values[0]))
        colum2.append(int(values[1]))

result = [colum1, colum2]

json_output = json.dumps(result, indent=4)

with open(args.output_file, "w") as output_file:
    output_file.write(json_output)

print(f"JSON output saved to {args.output_file}")