#!/usr/bin/env python3

import sys

def subpixels_to_pixels(num_subpixels):
  """
  This function calculates the number of pixels given the number of subpixels.

  Args:
      num_subpixels: The number of subpixels.

  Returns:
      The number of pixels (float).
  """

  # Assuming 16 subpixels per pixel
  subpixels_per_pixel = 16

  # Calculate the number of pixels
  num_pixels = num_subpixels / subpixels_per_pixel

  return num_pixels

def pixels_per_second(num_pixels, frame_rate):
  """
  This function calculates the pixels per second.

  Args:
      num_pixels: The number of pixels per frame.
      frame_rate: The video frame rate (FPS).

  Returns:
      The pixels per second (float).
  """

  pixels_per_second = num_pixels * frame_rate
  return pixels_per_second

# Get input from command line arguments (if any)
if len(sys.argv) > 1:
  try:
    num_subpixels = float(sys.argv[1])
  except ValueError:
    print("Error: Invalid input. Please enter a number.")
    sys.exit(1)
else:
  # If no arguments provided, prompt for input
  num_subpixels = float(input("Enter the number of subpixels per frame: "))

# Get frame rate (assuming it's known beforehand)
frame_rate = 60  # Replace with your actual frame rate

# Calculate number of pixels and pixels per second
num_pixels = subpixels_to_pixels(num_subpixels)
pixels_per_second = pixels_per_second(num_pixels, frame_rate)

# Print the results
print(f"{num_subpixels} subpixels is equivalent to approximately {num_pixels:.2f} pixels per frame")
print(f"This translates to {pixels_per_second:.2f} pixels per second at {frame_rate} FPS")


