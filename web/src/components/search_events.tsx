import React, { useEffect, useState } from 'react';

import { ApiClient } from '../api_client';
import { EventModel } from '../models/event';
import { EventItem } from './event_item';

type Props = {
  client: ApiClient;
};

export const SearchEvents = ({ client }: Props) => {
  const [fetching, setFetching] = useState(true);
  const [hasError, setHasError] = useState(false);
  const [events, setEvents] = useState<EventModel[]>([]);
  const [page, setPage] = useState(1);

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
        {!fetching && !hasError && <ul>
          {
            events.map(event =>
              <li key={event.id}>
                <EventItem event={event} />
              </li>
            )
          }
        </ul>}
        {hasError && <p>An error occured.</p>}
      </div>
      <div>
        <button disabled={page < 2} onClick={() => setPage(page - 1)}>Previous</button>
        <span>{page}</span>
        <button onClick={() => setPage(page + 1)}>Next</button>
      </div>
    </div>
  );
};
