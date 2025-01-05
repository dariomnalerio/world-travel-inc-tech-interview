import { useCallback, useState } from "react";
import { AuthContext } from "../hooks/use-auth";
import Cookies from "js-cookie";
interface AuthProviderProps {
  children: React.ReactNode;
}

const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [userId, setUserId] = useState<string | null>(null);

  const updateUserId = useCallback(
    (userId: string | null) => setUserId(userId),
    []
  );

  const logout = () => {
    Cookies.remove("auth_token");
    updateUserId(null);
  };

  return (
    <AuthContext.Provider value={{ userId, updateUserId, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthProvider };
