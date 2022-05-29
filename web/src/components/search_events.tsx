import React, { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { ApiClient } from '../api_client';
import { Pager } from './pager';
import { EventModel } from '../models/event';
import { EventItem } from './event_item';

type Props = {
  client: ApiClient;
};

export const SearchEvents = ({ client }: Props) => {
  const [fetching, setFetching] = useState(true);
  const [hasError, setHasError] = useState(false);
  const [events, setEvents] = useState<EventModel[]>([]);

  const navigate = useNavigate();
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const queryPage = parseInt(query.get("page") || "1");
  const page = (queryPage === NaN) ? 1 : queryPage;

  const onPageChange = (page: number) => {
    const newQuery = new URLSearchParams(query);
    newQuery.set("page", page.toString());
    navigate({ search: newQuery.toString() });
  };

  useEffect(() => {
    const fetch = async () => {
      try {
        const events = await client.getEvents(page);
        setEvents(events);
      } catch (ex) {
        console.error(ex);
        setHasError(true);
      }
      setFetching(false);
    };

    fetch();
  }, [page]);

  return (
    <div>
      <div>
        {fetching && <span>Loading...</span>}
        {!fetching && !hasError && <ul className='event-search-list'>
          {
            events.map(event =>
              <li key={event.id} className='event-search-list__event-item'>
                <EventItem event={event} />
              </li>
            )
          }
        </ul>}
        {hasError && <p>An error occured.</p>}
      </div>
      <Pager page={page} onPageChange={onPageChange} />
    </div>
  );
};
