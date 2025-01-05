import { createContext, use } from "react";

interface AuthContextProps {
  userId: string | null;
  updateUserId: (userId: string | null) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextProps | undefined>(
  undefined
);

const useAuth = () => {
  const context = use(AuthContext);

  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }

  return context;
};

export { useAuth };
