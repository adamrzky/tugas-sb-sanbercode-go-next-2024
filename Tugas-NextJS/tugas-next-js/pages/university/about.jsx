// pages/university/about.jsx
import Navbar from '../../components/Navbar';

const About = () => {
  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold">About Us</h1>
        <p className="mt-4">This is the about page of the University App.</p>
      </div>
    </>
  );
};

export default About;
