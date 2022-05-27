import React, { useState, useEffect } from 'react';
import { useLocation, useNavigate, useParams } from 'react-router-dom';

import { ApiClient } from '../api_client';
import { UserRank } from '../models/user_rank';
import { Pager } from './pager';

type Props = { client: ApiClient }

export const EventLeaderboard = ({ client }: Props) => {
  const { id } = useParams<{ id: string }>();

  if (id === undefined) {
    throw new Error("Event id was not given.")
  }

  const navigate = useNavigate();
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const queryPage = parseInt(query.get("page") || "1");
  const page = (queryPage === NaN) ? 1 : queryPage;

  const [fetching, setFetching] = useState(true);
  const [count, _] = useState(20);
  const [ranks, setRanks] = useState<UserRank[]>([]);

  const onPageChange = (page: number) => {
    const newQuery = new URLSearchParams(query);
    newQuery.set("page", page.toString());
    navigate({ search: newQuery.toString() });
  };

  useEffect(() => {
    const fetch = async () => {
      try {
        setFetching(true);
        const ranks = await client.getLeaderboard(id, page, count);
        setRanks(ranks);
      } catch {
        // setFailed(true);
      }
      setFetching(false);
    };

    fetch();
  }, [page, count]);

  return (
    <div>
      <div>
        {!fetching &&
          <ul>
            {
              ranks.map(rank =>
                <li key={rank.userId}>
                  <div>
                    {rank.rank}
                  </div>
                  <div>
                    {rank.userId}
                  </div>
                </li>
              )
            }
          </ul>
        }
      </div>
      <Pager page={page} onPageChange={onPageChange} />
    </div>
  )
};
