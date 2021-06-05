import React, { createContext, useContext, useEffect, useState } from "react";

export type FbAuthResponse = {
    accessToken :string; 
    expiresIn: string;
    signedRequest: string;
    userID: string;
}

export type FbLoginResp = {
    status: string;
    authResponse: FbAuthResponse;
}

export type LoginState = {
    loggedIn: boolean;
    fbUserId: string | null;
}

const FbLoginContext = createContext<LoginState | null>(null);

export const useFbLogin = function() {
    return useContext(FbLoginContext);
}


const FaceBookLoginProvider : React.FC = (props) => {
    const [loginState, setLoginState] = useState<LoginState>({loggedIn: false, fbUserId: null})

    useEffect(() => getLoginStatus(), [])

    const getLoginStatus = function() {
        (window as any).FB?.getLoginStatus(function(response : FbLoginResp) {
            handleAuthResponce(response);
        });
    }

    const handleAuthResponce = function(response: FbLoginResp) {
        if (response.status === "connected"){
            setLoginState({loggedIn: true, fbUserId: response.authResponse.userID})
        }
    }

    console.log(loginState)
    return (
        <div>
            {!loginState.loggedIn &&(
            <div 
                className="fb-login-button" 
                data-width="" 
                data-size="large" 
                data-button-type="continue_with" 
                data-layout="rounded" 
                data-auto-logout-link="false" 
                data-use-continue-as="false"
                data-onlogin="location.reload()"></div>                
            )}
            {loginState.loggedIn &&(
                <FbLoginContext.Provider value={loginState}>
                    <div>Logged in as userID = {loginState.fbUserId}</div>
                    {props.children}
                </FbLoginContext.Provider>
            )}

        </div>
    )
}

export default FaceBookLoginProvider
