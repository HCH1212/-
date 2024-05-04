namespace go userhoster

struct user{
    1: optional i8 identify=0;
    2: string name;
    3: string passwd;
    4: string token;
}

struct hoster{
    1: optional i8 identify=0;
    2: string name;
    3: string passwd;
    4: string token;
    5: string flag;
}

service userhoster {
    void Register(user u, hoster h);
    void Login(user u, hoster h);
    string SetToken(user u, hoster h);
}