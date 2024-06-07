import requests


image = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS8MojOVM7WCGn7EXGs5U-BEAS5ApGHZRba1g&s"
books = ["https://hlfpunkeyhenzvvpcipl.supabase.co/storage/v1/object/sign/school-of-elec/software/os/Lecture_2-_Processes.pdf?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJzY2hvb2wtb2YtZWxlYy9zb2Z0d2FyZS9vcy9MZWN0dXJlXzItX1Byb2Nlc3Nlcy5wZGYiLCJpYXQiOjE3MTc3MzYxNTEsImV4cCI6MTc0OTI3MjE1MX0.aJDEd_dwebceM2Df7t1FLzMq-QAdMj7Pu2-jldiZkiA&t=2024-06-07T04%3A55%3A51.024Z", "https://hlfpunkeyhenzvvpcipl.supabase.co/storage/v1/object/sign/school-of-elec/software/os/Lecture_3-_Threads.pdf?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJzY2hvb2wtb2YtZWxlYy9zb2Z0d2FyZS9vcy9MZWN0dXJlXzMtX1RocmVhZHMucGRmIiwiaWF0IjoxNzE3NzM2MTkzLCJleHAiOjE3NDkyNzIxOTN9.tcxiKPCjgbXUk1SrIpSW3qhW7e5W2lo__ByfixGLdy8&t=2024-06-07T04%3A56%3A32.912Z","https://hlfpunkeyhenzvvpcipl.supabase.co/storage/v1/object/sign/school-of-elec/software/os/Lecture_4_-_Process_Scheduling.pdf?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJzY2hvb2wtb2YtZWxlYy9zb2Z0d2FyZS9vcy9MZWN0dXJlXzRfLV9Qcm9jZXNzX1NjaGVkdWxpbmcucGRmIiwiaWF0IjoxNzE3NzM2Mjg0LCJleHAiOjE3NDkyNzIyODR9.5PJqeKH4FlUp-5t8zEFTEQQwRua-tcgTV0nqiX6iEI8&t=2024-06-07T04%3A58%3A03.887Z","https://hlfpunkeyhenzvvpcipl.supabase.co/storage/v1/object/sign/school-of-elec/software/os/Lecture_5-_Process_Synchonization_revised.pdf?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJzY2hvb2wtb2YtZWxlYy9zb2Z0d2FyZS9vcy9MZWN0dXJlXzUtX1Byb2Nlc3NfU3luY2hvbml6YXRpb25fcmV2aXNlZC5wZGYiLCJpYXQiOjE3MTc3MzYzNTMsImV4cCI6MTc0OTI3MjM1M30.QTSvyQIpozAH9Rx8UEwwqub6evPG01VoRZUvgteGh8M&t=2024-06-07T04%3A59%3A12.389Z","https://hlfpunkeyhenzvvpcipl.supabase.co/storage/v1/object/sign/school-of-elec/software/os/Lecture_6-_Deadlocks.pdf?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJzY2hvb2wtb2YtZWxlYy9zb2Z0d2FyZS9vcy9MZWN0dXJlXzYtX0RlYWRsb2Nrcy5wZGYiLCJpYXQiOjE3MTc3MzY0MjEsImV4cCI6MTc0OTI3MjQyMX0.P9LltS5zBhngB1km94LrWw9XzccB0lgVcUobMQhM7CQ&t=2024-06-07T05%3A00%3A20.999Z"]

url = "http://localhost:8000/books/2"
body = [
    {
    "name" : "processes",
    "description" : "chapter 2 of os",
    "file" : books[0],
    "image" : image
}, 
    {
    "name" : "Threads",
    "description" : "chapter 3 of os",
    "file" : books[1],
    "image" : image
},
    {
    "name" : "Preocess scheduling",
    "description" : "chapter 4 of os",
    "file" : books[2],
    "image" : image
},
{
    "name" : "process synchonization",
    "description" : "chapter 5 of os",
    "file" : books[3],
    "image" : image
},
{
    "name" : "Dead locks",
    "description" : "chapter 6 of os",
    "file" : books[4],
    "image" : image
}
]

for i in range(len(body)):
    response = requests.post(url=url, json=body[i])
    print(response.status_code)
    print(response.json())