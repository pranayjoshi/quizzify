import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useDropzone } from "react-dropzone";
import Navbar from "../Layout/navbar";
import { useSelector } from "react-redux";

const CreatePage = () => {
  const { getRootProps, getInputProps } = useDropzone();
  const [quiz_name, setQuizName] = useState("");
  const author = useSelector((state: any) => state.token);
    const name = useSelector((state: any) => state.name);
  const [content, setContent] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();
  const handleSubmit = async () => {
    setIsLoading(true);
    const quiz = { content, name, author, quiz_name };
    console.log(quiz);
    var data;
    try {
      const response = await fetch('http://localhost:8082/create_quiz', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(quiz),
      });
      data = await response.json();
      console.log(data);
      if (response.status === 200) {
        console.log(`Successfully created quiz: ${quiz_name}`);
      } else {
        console.error(`Error: ${response.status}`);
      }
    } catch (error) {
      console.error(error);
    }
    finally {
        setIsLoading(false); // Set loading to false when the request is finished

        navigate(`/show_quiz_id${data.quiz_id}`); // Navigate to the show quiz page
      }
  };
  if (isLoading) {
    return <div className="bg-inherit h-screen text-xl text-white">Loading...</div>; // Show loading spinner when isLoading is true
  }
  return (
    <div className=" bg-inherit">
      <Navbar />
      <div className="mx-16">
        
      <div className="bg-clip-text text-transparent bg-gradient-to-r from-green-400 to-blue-500 text-3xl font-bold ml-8 pt-8">Create Quiz</div>
      <div className=" mx-6"><input name="quiz_name" onChange={(e) =>setQuizName(e.target.value)} type="text" className="rounded-md px-4 mt-4 mr-16  py-2 border border-gray-300 focus:outline-none bg-gray-700 text-white w-full lg:text-xl " placeholder="Enter Quiz Name..." value={quiz_name}/></div>
      <div className="flex h-[70vh] pt-8 px-4 text-center justify-center">
        <div {...getRootProps()} className=" flex-1 border-dashed border-2 p-6 m-3">
          <input {...getInputProps()} />
          <p className="text-white ">Drag 'n' drop some files here, or click to select files</p>
        </div>
        <div className="flex-1 border-dashed border-2 p-6 m-3">
          <textarea className="w-full h-full bg-inherit text-white" placeholder="Paste text here" name="content"
  value={content} onChange={(e) => setContent(e.target.value)}></textarea>
        </div>
      </div>
      <div className="flex justify-between p-6">
        <Link to="/" className="text-blue-500 hover:text-blue-700 bg-white font-bold py-2 px-4 rounded-full border border-blue-400">
          Back
        </Link>
        <button className="text-white bg-gradient-to-r from-green-400 to-blue-500 rounded-full px-6 py-2" onClick={handleSubmit}>
          Submit
        </button>
      </div>
    </div></div>
  );
};

export default CreatePage;