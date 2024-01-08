import os
import json

# Directory containing all the deployment folders
root_dir = os.getcwd()
target_subdirs = ['sepolia', 'synapse_sepolia', 'optimism_sepolia', 'polygon_mumbai', 'arb_sepolia', 'scroll_sepolia']

# Iterate over each specified subdirectory
for target_subdir in target_subdirs:
    subdir_path = os.path.join(root_dir, target_subdir)
    if os.path.exists(subdir_path):
        print(f"{target_subdir}:")  # Print the subdirectory name
        for file in os.listdir(subdir_path):
            if file.endswith('.json'):
                file_path = os.path.join(subdir_path, file)
                with open(file_path, 'r') as json_file:
                    data = json.load(json_file)
                    address = data.get('address', 'Address not found')
                    print(f"{file.replace('.json', '')}: {address}")
        print()  # Print a newline for better separation
    else:
        print(f"Directory '{target_subdir}' not found in '{root_dir}'")
