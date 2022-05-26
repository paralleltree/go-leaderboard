import React from "react";

type Props = {
  page: number;
  maxPage?: number;
  onPageChange(page: number): void;
};

export const Pager = ({ page, maxPage, onPageChange }: Props) => {
  return (
    <div>
      <button onClick={() => onPageChange(page - 1)} disabled={page < 2}>Previous</button>
      <span>{page}</span>
      <button onClick={() => onPageChange(page + 1)} disabled={(maxPage !== undefined) && page >= maxPage}>Next</button>
    </div>
  );
};
