import React from 'react';

import { EventModel } from '../models/event';

type Props = {
  event: EventModel;
}

export const EventItem = ({ event }: Props) => {

  return (
    <div>
      <div>
        <span>{event.name}</span>
      </div>
      <div>
        <a href={`/events/${event.id}`}>Show</a>
      </div>
    </div>
  )
};
