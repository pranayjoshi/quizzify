import React from "react";
import { Link } from "react-router-dom";
import { useDropzone } from "react-dropzone";
import Navbar from "../Layout/navbar";

const CreatePage = () => {
  const { getRootProps, getInputProps } = useDropzone();

  return (
    <div className=" bg-inherit">
      <Navbar />
      <div className="mx-16">
      <div className="bg-clip-text text-transparent bg-gradient-to-r from-green-400 to-blue-500 text-3xl font-bold ml-8 pt-8">Create Quiz</div>
      <div className="flex h-[75vh] pt-8 px-4">
        <div {...getRootProps()} className=" flex-1 border-dashed border-2 p-6 m-3">
          <input {...getInputProps()} />
          <p className="text-white ">Drag 'n' drop some files here, or click to select files</p>
        </div>
        <div className="flex-1 border-dashed border-2 p-6 m-3">
          <textarea className="w-full h-full bg-inherit" placeholder="Paste text here"></textarea>
        </div>
      </div>
      <div className="flex justify-between p-6">
        <Link to="/" className="text-blue-500 hover:text-blue-700 bg-white font-bold py-2 px-4 rounded-full border border-blue-400">
          Back
        </Link>
        <button className="text-white bg-gradient-to-r from-green-400 to-blue-500 rounded-full px-6 py-2">
          Submit
        </button>
      </div>
    </div></div>
  );
};

export default CreatePage;