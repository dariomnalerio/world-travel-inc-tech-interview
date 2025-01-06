import { render, screen } from "@testing-library/react";
import { AuthProvider } from "../contexts/auth-provider";
import { ViewProvider } from "../contexts/view-provider";
import { DogCard } from "../components/Profile/dog-card";

describe("Profile", () => {
  describe("Dog-card", () => {
    const renderDogCard = (url: string, liked: boolean = true) => {
      return render(
        <AuthProvider>
          <ViewProvider initialView="profile">
            <DogCard url={url} liked={liked} />
          </ViewProvider>
        </AuthProvider>
      );
    };

    it("should render a dog card", () => {
      renderDogCard("https://dog-image.com", true);
      const dogCard = screen.getByTestId("dogCard");
      expect(dogCard).toBeInTheDocument();
    });

    it("should render a dog image", () => {
      renderDogCard("https://dog-image.com", true);
      const dogImage = screen.getByAltText("Dog");
      expect(dogImage).toBeInTheDocument();
    });

    it("should render a heart icon", () => {
      renderDogCard("https://dog-image.com", true);
      const heartIcon = screen.getByTestId("heartIcon");
      expect(heartIcon).toBeInTheDocument();
    });

    it("should render an unlike button", () => {
      renderDogCard("https://dog-image.com", true);
      const unlikeButton = screen.getByRole("button");
      expect(unlikeButton).toBeInTheDocument();
    });
  });
});
