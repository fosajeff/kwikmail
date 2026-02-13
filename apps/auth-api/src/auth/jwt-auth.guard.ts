import {
  CanActivate,
  ExecutionContext,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt/dist/jwt.service';

interface JwtPayload {
  [key: string]: unknown;
}

interface AuthenticatedRequest extends Request {
  user?: JwtPayload;
}

@Injectable()
export class JwtAuthGuard implements CanActivate {
  constructor(private readonly jwtService: JwtService) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const req = context.switchToHttp().getRequest<AuthenticatedRequest>();

    const auth = req.headers['authorization'] as string | undefined;
    if (!auth) {
      throw new UnauthorizedException('Missing authorization header');
    }

    const [type, token] = auth.split(' ');
    if (type !== 'Bearer' || !token) {
      throw new UnauthorizedException('Invalid authorization header');
    }
    try {
      const payload: JwtPayload = await this.jwtService.verifyAsync(token);

      // attach to request
      req.user = payload;

      return true;
    } catch {
      throw new UnauthorizedException('Invalid token');
    }
  }
}
