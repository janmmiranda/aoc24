import json
import argparse

parser = argparse.ArgumentParser(description="Convert a file into JSON")
parser.add_argument("input_file", help="Path to input file")
parser.add_argument("output_file", help="Path to output file")

args = parser.parse_args()

result = []
print(f"Txt file input reading from {args.input_file}")

with open(args.input_file, "r") as file:
    for line in file:
        currList = []
        vals = line.split()
        for val in vals:
            currList.append(int(val))
        result.append(currList)

json_output = json.dumps(result, indent=4)

with open(args.output_file, "w") as output_file:
    output_file.write(json_output)

print(f"JSON output saved to {args.output_file}")