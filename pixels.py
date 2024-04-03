#!/usr/bin/env python3

def pixels_per_frame(pixels_per_second, frame_rate):
  """
  This function calculates the number of pixels per frame.

  Args:
      pixels_per_second: The number of pixels processed per second.
      frame_rate: The video frame rate (FPS).

  Returns:
      The number of pixels per frame (float).
  """

  # Calculate the number of pixels per frame
  pixels_per_frame = pixels_per_second / frame_rate

  return pixels_per_frame

# Get input from command line arguments (if any)
import sys
if len(sys.argv) > 1:
  try:
    pixels_per_second = float(sys.argv[1])
  except ValueError:
    print("Error: Invalid input. Please enter a number.")
    sys.exit(1)
else:
  # If no arguments provided, prompt for input
  pixels_per_second = float(input("Enter the number of pixels per second: "))

# Get frame rate (assuming it's known beforehand)
frame_rate = 60  # Replace with your actual frame rate

# Calculate number of pixels per frame
pixels_per_frame = pixels_per_frame(pixels_per_second, frame_rate)

# Print the results
print(f"{pixels_per_second:.2f} pixels per second translates to approximately {pixels_per_frame:.2f} pixels per frame at {frame_rate} FPS")


