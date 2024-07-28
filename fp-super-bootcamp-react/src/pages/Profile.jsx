import React, { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import { UserContext } from '../contexts/UserContext';
import Swal from 'sweetalert2';
import './Profile.css';

const Profile = () => {
    const [profile, setProfile] = useState({ fullName: '', bio: '' });
    const { user, setUser } = useContext(UserContext);

    useEffect(() => {
        if (!user || !user.token) {
            Swal.fire("Please log in to view this page.");
            return;
        }

        const fetchProfile = async () => {
            try {
                const response = await axios.get('https://fp-super-bootcamp-go.vercel.app/user/profile', {
                    headers: { Authorization: `Bearer ${user.token}` }
                });
                setProfile({ fullName: response.data.FullName, bio: response.data.Bio });
            } catch (error) {
                console.error('Failed to fetch profile', error);
                Swal.fire({
                    icon: "error",
                    title: "Failed to Load",
                    text: "Could not load profile data. Please try again.",
                });
            }
        };

        fetchProfile();
    }, [user]);

    const handleUpdateProfile = async (e) => {
        e.preventDefault();
        if (!user || !user.token) {
            Swal.fire("Please log in to view this page.");
            return;
        }

        try {
            const response = await axios.put('https://fp-super-bootcamp-go.vercel.app/user/profile', {
                FullName: profile.fullName,
                Bio: profile.bio
            }, {
                headers: { Authorization: `Bearer ${user.token}` }
            });
            setProfile({ fullName: response.data.FullName, bio: response.data.Bio });
            Swal.fire({
                icon: "success",
                title: "Profile Updated",
                text: "Your profile has been updated successfully.",
            });
        } catch (error) {
            console.error('Failed to update profile', error);
            Swal.fire({
                icon: "error",
                title: "Failed to Update",
                text: "Failed to update profile. Please try again.",
            });
        }
    };

    return (
        <div className="profile-container">
            <form onSubmit={handleUpdateProfile} className="profile-form">
                <h2>Edit Profile</h2>
                <input
                    type="text"
                    value={profile.fullName}
                    onChange={(e) => setProfile({ ...profile, fullName: e.target.value })}
                    placeholder="Full Name"
                />
                <textarea
                    value={profile.bio}
                    onChange={(e) => setProfile({ ...profile, bio: e.target.value })}
                    placeholder="Bio"
                />
                <button type="submit">Update Profile</button>
            </form>
        </div>
    );
};

export default Profile;
