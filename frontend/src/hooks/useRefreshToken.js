import axios from '../api/axios';
import useAuth from './useAuth';

const useRefreshToken = () => {
    const { auth, setAuth } = useAuth();
    
    const refresh = async () => {
        try {
            let refreshToken = auth?.refreshToken;
            if(!!auth){
                refreshToken = localStorage.getItem("refresh_token")
            }

            const response = await axios.post('/users/refresh', {
                "refresh_token": refreshToken
            },
            {
                headers: { 'Content-Type': 'application/json' }
            });
            
            // Ensure you're updating the refresh token as well, if the server provides it
            setAuth(prev => ({
                ...prev,
                roles: response.data.roles,
                accessToken: response.data.accessToken,
                refreshToken: response.data.refreshToken,
            }));
            
            return response.data.accessToken;
        } catch (error) {
            console.error(error);
        }
    }

    return refresh;
};


export default useRefreshToken;
