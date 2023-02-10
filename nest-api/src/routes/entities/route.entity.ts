import { Prop, Schema, SchemaFactory, raw } from '@nestjs/mongoose';

export type RouteDocument = Route & Document;

@Schema()
export class Route {
  @Prop()
  title: string;

  @Prop(
    raw({
      lat: { type: Number },
      lgn: { type: Number },
    }),
  )
  startPosition: { lat: number; lng: number };

  @Prop(
    raw({
      lat: { type: Number },
      lgn: { type: Number },
    }),
  )
  endPosition: { lat: number; lng: number };
}

export const RouteSchema = SchemaFactory.createForClass(Route);
