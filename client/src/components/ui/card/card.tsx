import React, { JSX } from "react";
import styles from "./card.module.css";

type CardProps = React.ComponentPropsWithRef<"div"> & {
  children: React.ReactNode;
};

const Card = ({ children, ...props }: CardProps) => {
  return (
    <div data-testid="card" className={styles.card} {...props}>
      {children}
    </div>
  );
};

type HeadingLevel = 1 | 2 | 3 | 4 | 5 | 6;

type TitleProps<Level extends HeadingLevel> =
  React.ComponentPropsWithRef<`h${Level}`> & {
    children: React.ReactNode;
    headingLevel?: HeadingLevel;
  };

const Title = ({
  children,
  headingLevel = 2,
  ...props
}: TitleProps<HeadingLevel>) => {
  const Heading = `h${headingLevel}` as keyof Pick<
    JSX.IntrinsicElements,
    "h1" | "h2" | "h3" | "h4" | "h5" | "h6"
  >;

  return (
    <Heading className={styles.cardTitle} {...props}>
      {children}
    </Heading>
  );
};

type ContentProps = React.ComponentPropsWithRef<"div"> & {
  children: React.ReactNode;
};

const Content = ({ children, ...props }: ContentProps) => {
  return (
    <div data-testid="cardContent" className={styles.cardContent} {...props}>
      {children}
    </div>
  );
};

type FooterType = React.ComponentPropsWithRef<"div"> & {
  children: React.ReactNode;
};

const Footer = ({ children, ...props }: FooterType) => {
  return (
    <div data-testid="cardFooter" {...props}>
      {children}
    </div>
  );
};

Card.Title = Title;
Card.Content = Content;
Card.Footer = Footer;
export { Card };
