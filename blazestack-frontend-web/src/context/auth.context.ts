import { createContext } from 'react';

export interface IAuthUser {
  name?: string;
  email?: string;
  token?: string;
}

interface IAuthContext {
  user: IAuthUser;
  setUser: (user: IAuthUser) => void;
}

export const DefaultAuthContext: IAuthContext = {
  setUser: () => {},
  user: {},
};

export const AuthContext = createContext<IAuthContext>(DefaultAuthContext);
