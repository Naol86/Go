import os
from supabase import create_client, Client
from dotenv import load_dotenv

# Load environment variables from a .env file (optional)
load_dotenv()

SUPABASE_URL = "https://hlfpunkeyhenzvvpcipl.supabase.co"
SUPABASE_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImhsZnB1bmtleWhlbnp2dnBjaXBsIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTc3MDI3NDcsImV4cCI6MjAzMzI3ODc0N30.K1Xy_9mmuJR4fy00mk7W9pQp-fe_AKM46e6m31kLnJ4"

# Initialize the Supabase client
supabase: Client = create_client(SUPABASE_URL, SUPABASE_KEY)

def upload_file(bucket_name, file_path):
    with open(file_path, "rb") as file:
        file_content = file.read()

    file_name = os.path.basename(file_path)
    response = supabase.storage.from_(bucket_name).upload(file_name, file_content)
    
    if response.get("error"):
        print(f"Error uploading file: {response['error']}")
    else:
        print(f"File uploaded successfully: {response['data']}")

# Example usage
bucket_name = "bookstore"  # Replace with your bucket name
file_path = "/home/naol/Documents/program/Go/os4.pptx"  # Replace with the path to your file

upload_file(bucket_name, file_path)
