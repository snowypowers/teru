CREATE TABLE location (
	name VARCHAR(64) PRIMARY KEY,
	latitude DECIMAL(5,10),
	longitude DECIMAL(5,10),
	region VARCHAR(64)
);

CREATE TABLE wf2 (
	area VARCHAR(64),
	time TIMESTAMP,
	forecast VARCHAR(64) NOT NULL,
	FOREIGN KEY (area) REFERENCES location(name),
	CONSTRAINT area_time PRIMARY KEY (area, time)
);
