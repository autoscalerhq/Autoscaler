"use client"
import React, {createContext, useContext} from 'react';
import {JwtPayload} from 'jsonwebtoken';
export type AuthContextValue = {
    accessTokenPayload: JwtPayload;
};

export const AuthContext = createContext<AuthContextValue|null>(null);

export function useAuthContext() {
    const value =  useContext(AuthContext);
    if (value === null) {
        throw new Error('useAuthContext must be used within a AuthContext.Provider');
    }
    return value;
}


export function AuthContextProvider({value, children}: {value: AuthContextValue; children: React.ReactNode}) {
    return <AuthContext.Provider value={value}>
        {children}
    </AuthContext.Provider>
}