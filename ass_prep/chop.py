from PIL import Image

def chop_tiles(image_path, tile_size=(32, 32)):
  """
  Chops an image into tiles of specified size and saves them as separate PNGs.

  Args:
    image_path: Path to the image file.
    tile_size: A tuple representing the size (width, height) of each tile.
  """
  # Open the image
  img = Image.open(image_path).convert("RGBA")  # Ensure PNG format with alpha channel
  width, height = img.size

  # Calculate number of tiles
  num_x_tiles = width // tile_size[0]
  num_y_tiles = height // tile_size[1]

  # Loop through each tile and save
  for y in range(num_y_tiles):
    for x in range(num_x_tiles):
      # Define tile coordinates
      box = (x * tile_size[0], y * tile_size[1], (x + 1) * tile_size[0], (y + 1) * tile_size[1])
      # Crop the tile
      tile = img.crop(box)
      # Save the tile with filename based on position
      tile_filename = f"{image_path[:-4]}_{y}_{x}.png"  # Remove extension and add coordinates
      tile.save(tile_filename)

# Edit this path to your PNG file
image_path = "knight.png"

chop_tiles(image_path)

print("Image successfully chopped into tiles!")
