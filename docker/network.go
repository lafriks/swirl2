package docker

import (
	"context"
	"sort"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// NetworkCreate create a network.
func (d *Docker) NetworkCreate(ctx context.Context, name string, options *network.CreateOptions) error {
	return d.call(func(client *client.Client) error {
		resp, err := client.NetworkCreate(ctx, name, *options)
		if err == nil && resp.Warning != "" {
			d.logger.Warnf("network '%s' was created but got warning: %s", name, resp.Warning)
		}
		return err
	})
}

// NetworkList return all networks.
func (d *Docker) NetworkList(ctx context.Context) (networks []network.Summary, err error) {
	err = d.call(func(c *client.Client) (err error) {
		networks, err = c.NetworkList(ctx, network.ListOptions{})
		if err == nil {
			sort.Slice(networks, func(i, j int) bool {
				return networks[i].Name < networks[j].Name
			})
		}
		return
	})
	return
}

// NetworkCount return number of networks.
func (d *Docker) NetworkCount(ctx context.Context) (count int, err error) {
	err = d.call(func(c *client.Client) (err error) {
		var networks []network.Summary
		networks, err = c.NetworkList(ctx, network.ListOptions{})
		if err == nil {
			count = len(networks)
		}
		return
	})
	return
}

// NetworkRemove remove a network.
func (d *Docker) NetworkRemove(ctx context.Context, name string) error {
	return d.call(func(c *client.Client) (err error) {
		return c.NetworkRemove(ctx, name)
	})
}

// NetworkDisconnect Disconnect a container from a network.
func (d *Docker) NetworkDisconnect(ctx context.Context, network, container string) error {
	return d.call(func(c *client.Client) (err error) {
		return c.NetworkDisconnect(ctx, network, container, false)
	})
}

// NetworkInspect return network information.
func (d *Docker) NetworkInspect(ctx context.Context, name string) (n network.Summary, raw []byte, err error) {
	var c *client.Client
	if c, err = d.client(); err == nil {
		n, raw, err = c.NetworkInspectWithRaw(ctx, name, network.InspectOptions{})
	}
	return
}

// NetworkNames return network names by id list.
func (d *Docker) NetworkNames(ctx context.Context, ids ...string) (names map[string]string, err error) {
	var (
		c      *client.Client
		n      network.Summary
		lookup = func(id string) (network.Summary, error) {
			if c == nil {
				var err error
				c, err = d.client()
				if err != nil {
					return network.Summary{}, err
				}
			}
			return c.NetworkInspect(ctx, id, network.InspectOptions{})
		}
	)

	names = make(map[string]string)
	for _, id := range ids {
		name, ok := d.networks.Load(id)
		if ok {
			names[id] = name.(string)
		} else {
			n, err = lookup(id)
			if err != nil {
				return nil, err
			}
			names[id] = n.Name
			d.networks.Store(id, n.Name)
		}
	}
	return
}
