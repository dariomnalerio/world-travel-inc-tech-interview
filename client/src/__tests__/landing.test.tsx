import { render, screen } from "@testing-library/react";
import {
  CardSection,
  CardSectionProps,
} from "../components/Landing/card-section";

describe("Landing Card Section", () => {
  const renderCardSection = (customProps?: Partial<CardSectionProps>) => {
    const currentUrl =
      "https://images.dog.ceo/breeds/pyrenees/n02111500_7983.jpg";
    const fetchNextDog = vi.fn();
    const handleLike = vi.fn();
    const isFetching = false;
    const likeCurrentDog = false;

    const defaultProps = {
      currentUrl,
      fetchNextDog,
      handleLike,
      isFetching,
      likeCurrentDog,
    };

    const props = { ...defaultProps, ...customProps };
    render(<CardSection {...props} />);
  };

  it("should render the card section", () => {
    renderCardSection();
    const card = screen.getByTestId("card");
    expect(card).toBeInTheDocument();
  });

  it("should render the image", () => {
    renderCardSection();
    const img = screen.getByAltText("Random Dog");
    expect(img).toBeInTheDocument();
  });

  it("should render the like button", () => {
    renderCardSection();
    const heartIcon = screen.getByText("Like");
    expect(heartIcon).toBeInTheDocument();
  });

  it("should render the next dog button", () => {
    renderCardSection();
    const nextDogBtn = screen.getByText("Next");
    expect(nextDogBtn).toBeInTheDocument();
  });

  test("data-spin is true when fetching", () => {
    renderCardSection({ isFetching: true });
    const refreshIcon = screen.getByTestId("refreshIcon");
    expect(refreshIcon).toHaveAttribute("data-spin", "true");
  });

  test("data-spin is false when not fetching", () => {
    renderCardSection({ isFetching: false });
    const refreshIcon = screen.getByTestId("refreshIcon");
    expect(refreshIcon).toHaveAttribute("data-spin", "false");
  });

  test("next dog button should be disabled when fetching", () => {
    renderCardSection({ isFetching: true });
    const nextDogBtn = screen.getByText("Next");
    expect(nextDogBtn).toBeDisabled();
  });

  test("next dog button should be enabled when not fetching", () => {
    renderCardSection({ isFetching: false });
    const nextDogBtn = screen.getByText("Next");
    expect(nextDogBtn).not.toBeDisabled();
  });

  test("like button heart icon should have 100% opacity when current dog is liked", () => {
    renderCardSection({ likeCurrentDog: true });
    const heartIcon = screen.getByTestId("heartIcon");
    expect(heartIcon).toHaveAttribute("fill-opacity", "1");
  });

  test("like button heart icon should have 0% opacity when current dog is not liked", () => {
    renderCardSection({ likeCurrentDog: false });
    const heartIcon = screen.getByTestId("heartIcon");
    expect(heartIcon).toHaveAttribute("fill-opacity", "0");
  });

  test("like button should be enabled when not fetching", () => {
    renderCardSection({ isFetching: false });
    const likeBtn = screen.getByText("Like");
    expect(likeBtn).not.toBeDisabled();
  });

  test("like button should be disabled when fetching", () => {
    renderCardSection({ isFetching: true });
    const likeBtn = screen.getByText("Like");
    expect(likeBtn).toBeDisabled();
  });
});
