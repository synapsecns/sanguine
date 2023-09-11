# Ensure a directory is provided
if [[ -z "$1" ]]; then
    echo "Usage: $0 <directory>"
    exit 1
fi

# Recursively search and replace within the provided directory
find "$1" -type f -exec grep -l 'synapsecns/sanguine/blob/' {} + | while read -r file; do
    sed -i '/synapsecns\/sanguine\/blob\//s/contracts\//packages\/contracts-core\/contracts\//g' "$file"
done
