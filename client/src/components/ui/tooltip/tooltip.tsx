import { JSX } from "react";
import styles from "./tooltip.module.css";
interface TooltipProps {
  children: React.ReactNode;
  text: string;
}

const Tooltip = ({ children, text }: TooltipProps): JSX.Element => {
  return (
    <div className={styles.tooltip}>
      {children}
      {text.length > 0 && (
        <span data-testid="tooltipText" className={styles.tooltipText}>
          {text}
        </span>
      )}
    </div>
  );
};

export { Tooltip };
