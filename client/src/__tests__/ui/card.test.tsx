import { screen, render } from "@testing-library/react";
import { Card } from "../../components/ui/card/card";

describe("Card", () => {
  it("renders a div", () => {
    render(<Card>Card</Card>);
    const card = screen.getByTestId("card");
    expect(card).toBeInTheDocument();
  });

  it("renders children", () => {
    render(
      <Card>
        <h2>Card</h2>
      </Card>
    );
    const card = screen.getByRole("heading");
    expect(card).toBeInTheDocument();
  });

  it("renders custom props", () => {
    render(
      <Card className="card" data-custom="test">
        Card
      </Card>
    );
    const card = screen.getByTestId("card");
    expect(card).toHaveClass("card");
    expect(card).toHaveAttribute("data-custom", "test");
  });

  describe("Title", () => {
    it("renders a title", () => {
      render(<Card.Title>Title</Card.Title>);
      const title = screen.getByRole("heading");
      expect(title).toBeInTheDocument();
    });

    it("renders a title with a different heading level", () => {
      render(<Card.Title headingLevel={3}>Title</Card.Title>);
      const title = screen.getByRole("heading");
      expect(title.tagName).toBe("H3");
    });

    it("renders custom props", () => {
      render(
        <Card.Title className="title" data-custom="test">
          Title
        </Card.Title>
      );
      const title = screen.getByRole("heading");
      expect(title).toHaveClass("title");
      expect(title).toHaveAttribute("data-custom", "test");
    });

    it("renders a title as <Card> child", () => {
      render(
        <Card>
          <Card.Title>Title</Card.Title>
        </Card>
      );
      const title = screen.getByRole("heading");
      expect(title).toBeInTheDocument();
    });
  });

  describe("Content", () => {
    it("renders content", () => {
      render(<Card.Content>Content</Card.Content>);
      const content = screen.getByTestId("cardContent");
      expect(content).toBeInTheDocument();
    });

    it("renders children", () => {
      render(
        <Card.Content>
          <p>Content</p>
        </Card.Content>
      );
      const content = screen.getByRole("paragraph");
      expect(content).toBeInTheDocument();
    });

    it("renders custom props", () => {
      render(
        <Card.Content className="content" data-custom="test">
          Content
        </Card.Content>
      );
      const content = screen.getByTestId("cardContent");
      expect(content).toHaveClass("content");
      expect(content).toHaveAttribute("data-custom", "test");
    });

    it("renders content as <Card> child", () => {
      render(
        <Card>
          <Card.Content>Content</Card.Content>
        </Card>
      );
      const content = screen.getByTestId("cardContent");
      expect(content).toBeInTheDocument();
    });
  });

  describe("Footer", () => {
    it("renders a card footer", () => {
      render(<Card.Footer>Footer</Card.Footer>);
      const footer = screen.getByTestId("cardFooter");
      expect(footer).toBeInTheDocument();
    });

    it("renders children", () => {
      render(
        <Card.Footer>
          <p>Footer</p>
        </Card.Footer>
      );
      const footer = screen.getByRole("paragraph");
      expect(footer).toBeInTheDocument();
    });

    it("renders custom props", () => {
      render(
        <Card.Footer className="footer" data-custom="test">
          Footer
        </Card.Footer>
      );
      const footer = screen.getByTestId("cardFooter");
      expect(footer).toHaveClass("footer");
      expect(footer).toHaveAttribute("data-custom", "test");
    });

    it("renders footer as <Card> child", () => {
      render(
        <Card>
          <Card.Footer>Footer</Card.Footer>
        </Card>
      );
      const footer = screen.getByTestId("cardFooter");
      expect(footer).toBeInTheDocument();
    });
  });
});
